// import Looping from './Looping.jsx'
// import Mood from './Mood.jsx'
// import Button from './Button.jsx'
import Button2 from './Button2.jsx'
import { useCallback, useMemo, useState } from 'react'
import ChildComponent from './ChildComponend.jsx';

function App2() {
  const [likeCounter, setLikeCounter] = useState(0);
  const [subscriberCounter, setSubscriberCounter] = useState(0);
  const [name, setName] = useState('Lutfan');

  function pesanLike() {
    console.log(`Pesan Like rendered`);
    return (likeCounter < 10) ? 'like kurang dari 10' : 'like lebih dari 10';
  }
  // useMemo
  const displayLike = useMemo(() => pesanLike(), [likeCounter]);

  function handlerName() {
    let ChannelName = '';
    if (name === 'Lutfan') {
      ChannelName = "WPU";
    }else{
      ChannelName = "Lutfan";
    }

    setName(ChannelName);
    console.log(`Nama Channel ${name}`);
    
  }
  // useCallback(function, [dependencies]);
  const handlerName2 = useCallback(() => {
    let ChannelName = '';
    if (name === 'Lutfan') {
      ChannelName = "WPU";
    }else{
      ChannelName = "Lutfan";
    }

    setName(ChannelName);
    console.log(`Nama Channel ${name}`);
    
  }, [name])

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
      <ChildComponent name={name} aksi={handlerName2}/>
      
    </>
  )
}

export default App2