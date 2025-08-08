import { useEffect } from "react";
import { useState } from "react";
function Button2() {
    const [countLike, setCountLike] = useState(0);
    const [countDislike, setCountDislike] = useState(0);
    useEffect(() => {
        console.log(`Like ${countLike} || Dislike ${countDislike}`);
    }, [countLike]);

    return (
        <>
            <button onClick={() => setCountLike(countLike + 1)}>{countLike} Like</button>
            <button onClick={() => setCountDislike(countDislike + 1)}>{countDislike} Like</button>
        </>
    )
}

export default Button2