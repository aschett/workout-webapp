import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getWorkoutByID } from "../api/api-workout";
import type { Workout, WorkoutEntry } from "../types/workout";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";

function ListExercisesForWorkout() {
  const { id } = useParams<{ id: string }>();
  const [workout, setWorkout] = useState<Workout>();
  const [date, setDate] = useState<string>();


useEffect(() => {
  if (!id) return;

  getWorkoutByID(Number(id)).then((data) => {
    if (data) {
      setWorkout(workout);
      setDate(workout.date);
    }
  });
}, [id]);

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-semibold">Workout on {date}</h1>

      {!workout || workout.workouts.length === 0 ? (
        <p className="text-muted-foreground">No Workout entries for {date}</p>
      ) : (
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Exercise</TableHead>
              <TableHead>Sets</TableHead>
              <TableHead>Reps</TableHead>
              <TableHead>Weight (kg)</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {workout.workouts.map((entry, index) => (
              <TableRow key={index}>
                <TableCell>{entry.exerciseName}</TableCell>
                <TableCell>{entry.sets}</TableCell>
                <TableCell>{entry.reps}</TableCell>
                <TableCell>{entry.weight}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </div>
  );
}

export default ListExercisesForWorkout;
