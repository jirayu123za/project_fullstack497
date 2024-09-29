import AssignmentSubmit from "./AssignmentSubmit";

interface AssignmentSubmittedProps {
  ConfigAssignment: { StdCode: string; Status: string }[];
}

export default function AssignmentSubmitted({
  ConfigAssignment,
}: AssignmentSubmittedProps) {
  return (
    <div className="bg-white border-2 border-B1 w-[200px] max-h-[300px] rounded-xl p-6 flex flex-col space-y-5 overflow-y-auto">
      {ConfigAssignment.map((config, index) => (
        <AssignmentSubmit
          key={index}
          StdCode={config.StdCode}
          Status={config.Status}
        />
      ))}
    </div>
  );
}
