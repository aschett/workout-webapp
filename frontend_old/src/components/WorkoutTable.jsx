function WorkoutTable({ workouts }) {
  if (workouts.length === 0) {
    return <p>No exercises in this workout.</p>;
  }

  return (
    <div className="container">
      <table>
        <thead>
          <tr>
            <th>Exercise</th>
            <th>Weight (kg)</th>
            <th>Sets</th>
            <th>Reps</th>
          </tr>
        </thead>
        <tbody>
          {workouts.map((w) => (
            <tr key={w.id}>
              <td>{w.exerciseName}</td>
              <td>{w.weight}</td>
              <td>{w.sets}</td>
              <td>{w.reps}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default WorkoutTable;
