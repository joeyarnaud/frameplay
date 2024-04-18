import { useState } from 'react'
import ConvertDemo from './components/convert/ConvertDemo'
import InputDemo from './components/input/InputDemo'

function App() {
  const [isConvertDemo, setIsConvertDemo] = useState(false)
  return (
    <div>
      <header>Frameplay {isConvertDemo ? 'Convert' : 'Input'} Demo</header>
      <div>
        <button onClick={() => setIsConvertDemo(!isConvertDemo)}>
          View {!isConvertDemo ? 'Convert' : 'Input'} Demo
        </button>
      </div>
      {isConvertDemo ? <ConvertDemo /> : <InputDemo />}
    </div>
  )
}

export default App
