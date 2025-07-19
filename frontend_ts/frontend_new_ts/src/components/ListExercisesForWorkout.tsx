import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getWorkoutByID } from "../api/api-workout";
import type { WorkoutEntry } from "../types/workout";
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
  const [entries, setEntries] = useState<WorkoutEntry[]>([]);

  useEffect(() => {
    if (!id) return;
    getWorkoutByID(Number(id)).then(setEntries);
  }, [id]);

  return (
    <div className="space-y-6">
      <h1 className="text-2xl font-semibold">Workout {id}</h1>

      {entries.length === 0 ? (
        <p className="text-muted-foreground">No Workout entries</p>
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
            {entries.map((entry, index) => (
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
