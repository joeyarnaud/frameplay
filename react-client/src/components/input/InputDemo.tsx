import { useState } from 'react'

const inpStyle = {}

const labelStyle = {
  display: 'block',
}

const requestOptions = {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
}

const InputDemo = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const submit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const reqBody = { email, password }

    fetch('http://localhost:8080/', {
      ...requestOptions,
      body: JSON.stringify(reqBody),
    })
      .then((response) => response.json())
      .then((data) => console.log('Success:', data))
      .catch((error) => console.error('Error:', error))
  }

  return (
    <form onSubmit={submit}>
      <div>
        <label style={labelStyle}>email</label>
        <input type="email" onChange={(e) => setEmail(e.target.value)} />
      </div>
      <div>
        <label style={labelStyle}>password</label>
        <input type="password" onChange={(e) => setPassword(e.target.value)} />
      </div>
      <button type="submit">button</button>
    </form>
  )
}

export default InputDemo
