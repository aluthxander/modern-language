// import Looping from './Looping.jsx'
// import Mood from './Mood.jsx'
// import Button from './Button.jsx'
import Button2 from './Button2.jsx'
import { useMemo, useState } from 'react'

function App2() {
  const [likeCounter, setLikeCounter] = useState(0);
  const [subscriberCounter, setSubscriberCounter] = useState(0);

  function pesanLike() {
    console.log(`Pesan Like rendered`);
    return (likeCounter < 10) ? 'like kurang dari 10' : 'like lebih dari 10';
  }
  // useMemo
  const displayLike = useMemo(() => pesanLike(), [likeCounter]);

  return (
    <>
      <Button2 />
      <hr />
      <p>
        <button onClick={() => setLikeCounter(likeCounter + 1)}>{likeCounter} <b>Like</b></button> {displayLike}
      </p>
      <p>
        <button onClick={() => setSubscriberCounter(subscriberCounter + 1)}>{subscriberCounter} <b>Like</b></button>

      </p>
    </>
  )
}

export default App2