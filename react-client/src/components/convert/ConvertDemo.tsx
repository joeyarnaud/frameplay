import { useState } from 'react'
import { convertToObject } from '../../utils/convertToObject'

const ConvertDemo = () => {
  const [typeString, setTypeString] = useState('')

  return (
    <div>
      <textarea
        style={{ height: 200, width: 500 }}
        onChange={(e) => setTypeString(e.target.value)}
      ></textarea>
      <div>{JSON.stringify(convertToObject(typeString), null, 4)}</div>
    </div>
  )
}

export default ConvertDemo
