import AssignmentSubmit from "./AssignmentSubmit";

interface AssignmentSubmittedProps {
  students: { StdCode: string; Status: string }[]; // รับข้อมูลนักศึกษาจาก backend
}

export default function AssignmentSubmitted({ students }: AssignmentSubmittedProps) {
  return (
    <div className="bg-white border-2 border-B1 w-[200px] h-[500px] rounded-xl p-6 flex flex-col space-y-5 overflow-y-auto">
      {students.map((student, index) => (
        <AssignmentSubmit
          key={index}
          StdCode={student.StdCode}
          Status={student.Status} // "submitted" หรือ "not_submitted"
        />
      ))}
    </div>
  );
}
