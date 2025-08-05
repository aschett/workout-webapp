import { useEffect, useState } from "react";
import { getAllWorkouts } from "../api/api-workout";
import { createWorkout } from "../api/api-workout";
import { Link } from "react-router-dom";
import type { Workout } from "../types/workout";
import { Button } from "@/components/ui/button";
import { ArrowUpDown, Plus } from "lucide-react";
import { Calendar } from "@/components/ui/calendar";
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover";
import { format } from "date-fns";
import { Alert, AlertDescription, AlertTitle,} from "@/components/ui/alert";
import { AlertCircleIcon } from "lucide-react";
import { toast } from "sonner"



function ListAllWorkouts() {
  const [workouts, setWorkouts] = useState<Workout[]>([]);
  const [sortAsc, setSortAsc] = useState(true);
  const [adding, setAdding] = useState(false);
  const [selectedDate, setSelectedDate] = useState<Date | undefined>();
  const [error, setError] = useState<string | null>(null);

  const fetchWorkouts = async () => {
    const data = await getAllWorkouts();
    setWorkouts(data);
  };

  useEffect(() => {
    fetchWorkouts();
  }, []);

const handleDateSelect = async (date: Date | undefined) => {
  if (!date) return;
  setAdding(true);
  setError(null);
  try {
    await createWorkout({ date: format(date, "yyyy-MM-dd") });
    await fetchWorkouts();
  } catch (err: any) {
    console.error("Failed to add workout", err);
    setError(err.message);
  } finally {
    setAdding(false);
    setSelectedDate(undefined);
  }
};


  const sorted = [...workouts].sort((a, b) =>
    sortAsc ? a.date.localeCompare(b.date) : b.date.localeCompare(a.date)
  );


  useEffect(() => {
  if (error) {
    toast.custom(() => (
      <Alert variant="destructive" className="w-full max-w-sm">
        <AlertCircleIcon className="h-4 w-4" />
        <AlertTitle>Could not add workout</AlertTitle>
        <AlertDescription>{error}</AlertDescription>
      </Alert>
    ));
  }
}, [error]);


  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <h1 className="text-xl font-semibold">Workouts</h1>

        <div className="flex gap-2">
          <Popover>
            <PopoverTrigger asChild>
              <Button variant="outline" size="sm" className="flex items-center gap-1">
                <Plus className="w-4 h-4" />
                Add Workout
              </Button>
            </PopoverTrigger>
            <PopoverContent className="w-auto p-0">
              <Calendar
                mode="single"
                selected={selectedDate}
                onSelect={handleDateSelect}
                disabled={adding}
              />
            </PopoverContent>
          </Popover>

          <Button
            variant="ghost"
            onClick={() => setSortAsc((prev) => !prev)}
            className="flex items-center gap-2"
          >
            <ArrowUpDown className="w-4 h-4" />
            {sortAsc ? "Newest First" : "Oldest First"}
          </Button>
        </div>
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
