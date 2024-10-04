
interface AssignmentDetailProps {
  user_group_name: string;
}

export default function AssignmentDetail({ user_group_name }: AssignmentDetailProps) {
  return (
    <textarea
      className="container border-2 border-B1 min-w-full min-h-[500px] rounded-lg p-6 text-M1"
      placeholder="Assignment details..."
      disabled={user_group_name === "STUDENT" || user_group_name === ""}
    />
  );
}
