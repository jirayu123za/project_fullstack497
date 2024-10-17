import AssignmentSubmit from "./AssignmentSubmit";

interface Submission {
  user_id: string;
  user_name: string;
  user_submitted: boolean;
}

interface AssignmentSubmittedProps {
  submissions: Submission[];
}

export default function AssignmentSubmitted({ submissions }: AssignmentSubmittedProps) {
  const students = submissions.map((submission) => ({
    user_name: submission.user_name,
    user_submitted: submission.user_submitted ? true : false,
  }));

  return (
    <div className="bg-white border-2 border-B1 w-[270px] h-[500px] rounded-xl p-6 flex flex-col space-y-5 overflow-y-auto">
      {students.map((student, index) => (
        <AssignmentSubmit
          key={index}
          user_name={student.user_name}
          user_submitted={student.user_submitted}
        />
      ))}
    </div>
  );
}
