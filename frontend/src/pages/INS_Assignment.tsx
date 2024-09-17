import LeftMain from "../components/LeftMain";
import RightMain from '../components/RightMain';
import AssignmentDetail from '../components/AssignmentDetail';
import DropBox from '../components/DropBox';
import AssignmentSubmitted from '../components/AssignmentSubmitted';
import noticon from "../icons/bxs_bell.png";
import joinicon from "../icons/material-symbols_join.png";
import dashicon from "../icons/mdi_human-welcome.png";
import exiticon from "../icons/vaadin_exit-o.png";
import Assign from "../icons/ion_list.png";
import { useState } from 'react';

export default function InstructorDashboard() {
    const icons = [dashicon, noticon, joinicon, exiticon];
    const links = ["/dashboard", "/notifications", "/join", "/exit"];

    const [uploadedFiles, setUploadedFiles] = useState<string[]>([]);

    const handleUpdate = () => {
        console.log("Update Assignment Details");
    };

    const handleDelete = () => {
        console.log("Delete Assignment");
    };

    const handleFileUpload = (file: File) => {
        setUploadedFiles((prevFiles) => [...prevFiles, file.name]);
        console.log(`Uploaded: ${file.name}`);
    };

    const ConfigAssignment = [
        { StdCode: "640610629", Status: "#E61616" },
        { StdCode: "640629042", Status: "#E61616" },
        { StdCode: "633934788", Status: "#E61616" },
        { StdCode: "634124894", Status: "#E61616" },
        { StdCode: "123408895", Status: "#E61616" },
        { StdCode: "640610629", Status: "#E61616" },
        { StdCode: "640629042", Status: "#E61616" },
        { StdCode: "633934788", Status: "#E61616" },
        { StdCode: "634124894", Status: "#E61616" },
        { StdCode: "123408895", Status: "#E61616" },
        { StdCode: "640610629", Status: "#E61616" },
        { StdCode: "640629042", Status: "#E61616" },
        { StdCode: "633934788", Status: "#E61616" },
        { StdCode: "634124894", Status: "#E61616" },
        { StdCode: "123408895", Status: "#E61616" },
    ];

    return (
        <div className="bg-B1 h-screen flex justify-center items-center gap-8">
            <div className="w-[1174px] h-[900px] bg-white rounded-2xl relative  ">
                <LeftMain icon={Assign} title="Assignment 1" />
                <div className="absolute top-[100px] left-0 right-0 px-10 flex justify-between items-center mb-2 mt-[10px] ">
                    <p className="text-[16px] font-semibold text-[#344B59]">Due: 09-Aug-2024</p>
                    <div className="flex items-center space-x-2">
                        <button
                            onClick={handleUpdate}
                            className="bg-yellow-300 text-white px-4 py-2 rounded-full text-sm w-[125px]"
                        >
                            Update
                        </button>
                        <button
                            onClick={handleDelete}
                            className="bg-red-500 text-white px-4 py-2 rounded-full text-sm w-[125px]"
                        >
                            Delete
                        </button>
                    </div>
                </div>

                <div className="absolute  top-[-60px] left-10 flex items-start gap-5 ">
                    <div className="">
                        <AssignmentDetail />
                    </div>
                    <div className="">
                        <AssignmentSubmitted ConfigAssignment={ConfigAssignment} />
                    </div>
                </div>

                <div className="absolute bottom-[30px] left-10 flex items-start">
                    <DropBox onFileUpload={handleFileUpload} />
                    <div className="flex flex-col ml-4 space-y-2">
                        {uploadedFiles.map((fileName, index) => (
                            <div
                                key={index}
                                className="flex items-center border-2 border-[#5A8FAA] rounded-lg p-2 w-[214px]"
                            >
                                <img
                                    src="/icons/attachment_icon.png"
                                    alt="attachment"
                                    className="w-4 h-4 mr-2"
                                />
                                <p className="text-[#5A8FAA]">{fileName}</p>
                            </div>
                        ))}
                    </div>
                </div>
            </div>

            {/* RightMain Component */}
            <RightMain icons={icons} links={links} />
        </div>
    );
}