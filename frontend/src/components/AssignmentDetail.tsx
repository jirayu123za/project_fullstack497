import { useState } from "react";

interface AssignmentDetailProps {
  user_group_name: string;
  assignment_description: string;
}

export default function AssignmentDetail({ user_group_name, assignment_description }: AssignmentDetailProps) {
  const [isFocused, setIsFocused] = useState(false);

  const handleFocus = () => {
    setIsFocused(true);
  };

  const handleBlur = (e: React.FocusEvent<HTMLTextAreaElement>) => {
    if (e.target.value.trim() === "") {
      setIsFocused(false);
    }
  };


  return (
    <textarea
      className="container border-2 border-B1 min-w-full min-h-[500px] rounded-lg p-6 text-M1"
      placeholder={!isFocused ? "Assignment details..." : ""}
      defaultValue={isFocused ? assignment_description : assignment_description}
      readOnly={user_group_name === "STUDENT" || user_group_name === ""}
      onFocus={handleFocus}
      onBlur={handleBlur}
    />
  );
}
