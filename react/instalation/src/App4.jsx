import { useRef } from "react"

function App() {
  const pesan = useRef(null)
  const displayPesan = useRef(null)

  function clickHandler() {
    console.log(pesan.current.value);
    displayPesan.current.innerHTML = pesan.current.value;
    displayPesan.current.style.color = "blue";
    displayPesan.current.style.padding = "4px";
    displayPesan.current.style.border = "1px solid black";
    displayPesan.current.style.borderRadius = "4px";
  }
  return (
    <>
      <div>
        <input type="text" ref={pesan} placeholder="Kirim pesan kamu" />
      </div>
      <div>
        <button onClick={clickHandler}>Klik Aku</button>
      </div>
      <div ref={displayPesan}></div>
    </>
  )
}

export default App