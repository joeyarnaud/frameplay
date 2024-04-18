import * as ts from "typescript";

// Function to visit nodes in the type definition and convert to a JavaScript object
const visit = (node: ts.Node): any => {
    // console.log('NODE:::', node);
    switch (node.kind) {
        case ts.SyntaxKind.TypeLiteral:
            let obj: any = {};
            ts.forEachChild(node, (child) => {
                if (ts.isPropertySignature(child) && child.name && child.type) {
                    const key = (child.name as ts.Identifier).text;
                    obj[key] = visit(child.type);
                }
            });
            return obj;

        case ts.SyntaxKind.UnionType:
            const union: ts.UnionTypeNode = node as ts.UnionTypeNode;

            return union.types.map(type => {
                return visit(type)
            });

        case ts.SyntaxKind.StringLiteral:
            return (node as ts.StringLiteral).text;

        case ts.SyntaxKind.NumberKeyword:
            return 'number';

        case ts.SyntaxKind.BooleanKeyword:
            return 'boolean';

        case ts.SyntaxKind.StringKeyword:
            return 'string';

        default:
            return node.getText().replace(/["']/g, '');
    }
}

// Function to convert a TypeScript type string to an object representation
const convertToObject = (typeString: string): any =>  {
    const [typeName] = typeString.replace('type', '').split('=').map(o => o.trim());

    // Prepare the source file
    const sourceFile = ts.createSourceFile(
        "temp.ts",
        typeString,
        ts.ScriptTarget.Latest,
        true,
        ts.ScriptKind.TS
    );
    
    // Find the TypeAliasDeclaration and convert
    let result: any = {};
    ts.forEachChild(sourceFile, node => {
        if (ts.isTypeAliasDeclaration(node) && node.type) {
            result = visit(node.type);
        }
    });

    return {
        [typeName.toLowerCase()]: {
            ...result
        }
    };
}

export { convertToObject }

// Example usage
const typeString = `type Button = {
    variant: "solid" | "text";
    thing: string;
    thing2: {
        variant: 'big' | 'small';
    };
    thing3: boolean;
};`;

console.log(convertToObject(typeString).button.thing2);