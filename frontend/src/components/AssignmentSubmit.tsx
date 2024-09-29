interface AssignmentSubmitProps {
  StdCode: string;
  Status: string; // "submitted" or "not_submitted"
}

export default function AssignmentSubmit({ StdCode, Status }: AssignmentSubmitProps) {
  // กำหนดสีวงกลมตามสถานะการส่งงาน
  const circleColor = Status === "submitted" ? "#93B955" : "#E61616"; // เขียวสำหรับส่งงานแล้ว, แดงสำหรับยังไม่ส่ง

  return (
    <div className="flex justify-start items-center gap-3 p-2">
      <div>
        {/* วงกลมแสดงสถานะ */}
        <svg width="29" height="29" viewBox="0 0 29 29" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="14.5" cy="14.5" r="14.5" fill={circleColor} />
        </svg>
      </div>
      <div>{StdCode}</div> {/* รหัสนักศึกษา */}
    </div>
  );
}
