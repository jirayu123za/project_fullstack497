import React, { useCallback } from "react";
import { useDropzone, Accept } from "react-dropzone";
import { FaCloudUploadAlt } from "react-icons/fa";
import axios from "axios";

interface DropBoxProps {
  onFileUpload: (file: File) => void;
}

const DropBox: React.FC<DropBoxProps> = ({ onFileUpload }) => {
  const onDrop = useCallback(
    async (acceptedFiles: File[]) => {
      acceptedFiles.forEach(async (file) => {
        if (file.type === "application/pdf") {
          // Call the file upload handler function
          onFileUpload(file);

          // Prepare form data for API request
          const formData = new FormData();
          formData.append("file", file);

          try {
            // Send the file to the backend API using axios
            const response = await axios.post("https://your-backend-api/upload", formData, {
              headers: {
                "Content-Type": "multipart/form-data",
              },
            });

            // Handle response
            if (response.status === 200) {
              console.log("File uploaded successfully:", response.data);
            } else {
              console.error("Failed to upload file:", response.data);
            }
          } catch (error) {
            console.error("Error uploading file:", error);
          }
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
  });

  return (
    <div
      {...getRootProps()}
      className="flex flex-col items-center justify-center w-full h-44 border-2 border-dashed border-blue-200 rounded-lg cursor-pointer"
    >
      <input {...getInputProps()} />
      <div className="flex flex-col items-center justify-center">
        <FaCloudUploadAlt className="h-12 w-12 text-blue-300" />
        <p className="mt-2 text-gray-500">
          Drag or <span className="text-blue-500 cursor-pointer">Browse</span>{" "}
          files to upload
        </p>
      </div>
    </div>
  );
};

export default DropBox;
