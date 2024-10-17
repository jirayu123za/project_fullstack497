interface RightMainProps {
  icons: string[];
  links: string[];
  profile_image: string;
  user_group_name: string;
  handleLogout: () => void; 
}

export default function RightMain({
  icons,
  links,
  profile_image,
  user_group_name,
  handleLogout,
}: RightMainProps) {

  return (
    <div className="bg-white rounded-xl w-[150px] h-full">
      <div className="flex flex-col items-center h-full justify-between py-6">
        <div className="flex flex-col justify-center items-center gap-10 ">
          <div>
            <img
              className="w-28 h-28 p-2 rounded-full ring-4 ring-B1 object-cover"
              src={profile_image}
              alt="Profile image"
            />
          </div>

          <div className="flex flex-col gap-10 w-fit">
            {icons.map((icon, index) => (
              <a href={links[index]} key={index} className="inline-block">
                <img
                  src={icon}
                  alt={`icon-${index}`}
                  className={`w-14 h-14 inline-block ${index === icons.length - 1 ? 'cursor-pointer' : ''}`}
                  onClick={index === icons.length - 1 ? handleLogout : undefined}
                />
              </a>
            ))}
          </div>
        </div>
        {/* Role button */}
        <button
          className={`text-white font-medium bg-E1 px-4 py-2 rounded-full w-[130px] ${user_group_name.length > 6 ? 'text-md' : 'text-xl'}`}
          disabled
        >
          {user_group_name}
        </button>
      </div>
    </div>
  );
}