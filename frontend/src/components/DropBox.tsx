import React, { useCallback } from "react";
import { useDropzone, Accept } from "react-dropzone";
import { FaCloudUploadAlt } from "react-icons/fa";

interface DropBoxProps {
  onFileUpload: (file: File) => void;
  isDuePassed: boolean;
}

const DropBox: React.FC<DropBoxProps> = ({ onFileUpload, isDuePassed }) => {
  const onDrop = useCallback(
    async (acceptedFiles: File[]) => {
      acceptedFiles.forEach((file) => {
        if (file.type === "application/pdf") {
          onFileUpload(file);
        } else {
          alert("Only PDF files are allowed!");
        }
      });
    },
    [onFileUpload]
  );

  const accept: Accept = { "application/pdf": [] };


  const { getRootProps, getInputProps } = useDropzone({
    onDrop,
    accept,
    disabled: isDuePassed,
  });

  return (
    <div
      {...getRootProps()}
      className={`flex flex-col items-center justify-center w-full h-44 border-2 border-dashed ${
        isDuePassed ? "border-gray-300 cursor-not-allowed" : "border-blue-200 cursor-pointer"
      } rounded-lg`}
    >
      <input {...getInputProps()} disabled={isDuePassed} />
      <div className="flex flex-col items-center justify-center">
        <FaCloudUploadAlt className={`h-12 w-12 ${isDuePassed ? "text-gray-300" : "text-blue-300"}`} />
        <p className={`mt-2 ${isDuePassed ? "text-gray-300" : "text-gray-500"}`}>
          {isDuePassed ? "Upload Disabled (Due Date Passed)" : "Drag or Browse files to upload"}
        </p>
      </div>
    </div>
  );
};

export default DropBox;
