import { useEffect, useState } from 'react';
import { getAllWorkouts } from '../api/api-workout';
import { Link } from 'react-router-dom';

function WorkoutList() {
  const [workouts, setWorkouts] = useState([]);

  useEffect(() => {
    getAllWorkouts().then(setWorkouts);
  }, []);

  return (
    <div className="container">
      <h1>Workouts</h1>
      <ul>
        {workouts.map((w) => (
          <li key={w.id}>
            <Link to={`/workouts/${w.id}`}>
              Workout from {w.date}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default WorkoutList;
