interface AssignmentSubmitProps {
  user_name: string;
  user_submitted: boolean;
}

export default function AssignmentSubmit({ user_name, user_submitted }: AssignmentSubmitProps) {
  const circleColor = user_submitted === true ? "#93B955" : "#E61616";

  return (
    <div className="flex justify-start items-center gap-3 p-2">
      <div>
        {/* วงกลมแสดงสถานะ */}
        <svg width="19" height="19" viewBox="0 0 29 29" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="14.5" cy="14.5" r="14.5" fill={circleColor} />
        </svg>
      </div>
      <div>{user_name}</div> {/* รหัสนักศึกษา */}
    </div>
  );
}
