import { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getWorkoutByID } from '../api/api-workout';
import WorkoutTable from './WorkoutTable';

function WorkoutPage() {
  const { id } = useParams();
  const [entries, setEntries] = useState([]);

  useEffect(() => {
    getWorkoutByID(id).then(setEntries);
  }, [id]);

  return (
    <div className="container">
      <h1>Workout {id}</h1>
      <WorkoutTable workouts={entries} />
    </div>
  );
}

export default WorkoutPage;
