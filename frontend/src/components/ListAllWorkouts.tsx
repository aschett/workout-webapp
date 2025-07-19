import { useEffect, useState } from "react";
import { getAllWorkouts } from "../api/api-workout";
import { Link } from "react-router-dom";
import type { Workout } from "../types/workout";
import { Button } from "@/components/ui/button";
import { ArrowUpDown } from "lucide-react";

function ListAllWorkouts() {
  const [workouts, setWorkouts] = useState<Workout[]>([]);
  const [sortAsc, setSortAsc] = useState(true);

  useEffect(() => {
    getAllWorkouts().then(setWorkouts);
  }, []);

  const sorted = [...workouts].sort((a, b) =>
    sortAsc
      ? a.date.localeCompare(b.date)
      : b.date.localeCompare(a.date)
  );

  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h1 className="text-xl font-semibold">Workouts</h1>
        <Button
          variant="ghost"
          onClick={() => setSortAsc((prev) => !prev)}
          className="flex items-center gap-2"
        >
          <ArrowUpDown className="w-4 h-4" />
          {sortAsc ? "Newest First" : "Oldest First"}
        </Button>
      </div>

      <ul className="space-y-1">
        {sorted.map((w) => (
          <li key={w.id} className="flex justify-between border-b py-2 text-sm">
            <span>{w.date}</span>
            <Link to={`/workouts/${w.id}`} className="text-primary underline">
              Inspect Details
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ListAllWorkouts;
