import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import ConvertDemo from './ConvertDemo'

describe('ConvertDemo Component', () => {
  test('should parse type definition and display correct output', () => {
    render(<ConvertDemo />)

    const expectedOutput = `{ "button": { "variant": [ "solid", "text" ], "thing": "string", "thing2": { "variant": [ "big", "small" ] }, "thing3": "boolean" } }`

    const textarea = screen.getByRole('textbox')
    const typeInput = `type Button = {
            variant: "solid" | "text";
            thing: string;
            thing2: {
                variant: 'big' | 'small';
            };
            thing3: boolean;
        };`

    fireEvent.change(textarea, { target: { value: typeInput } })

    const outputDiv = screen.getByText(expectedOutput)
    expect(outputDiv).toBeInTheDocument()
  })
})
