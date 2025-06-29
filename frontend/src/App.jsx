import { useEffect, useState } from 'react'

function App() {
  const [workouts, setWorkouts] = useState([])

  useEffect(() => {
    fetch('/api/workouts')
      .then(res => res.json())
      .then(data => setWorkouts(data))
  }, [])

  return (
    <ul>
      {workouts.map(w => (
        <li key={w.id}>
          {w.date}
        </li>
      ))}
    </ul>
  )
}

export default App
