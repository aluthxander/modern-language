import { useState } from "react"

function Button() {
    const [counter, setCounter] = useState(0);
    let newCounter = 0;
    function clickHandlerParams(nama) {
        console.log(`saya diklik oleh ${nama}`);
    }

    function clickHandler() {
        console.log('saya diklik');
    }

    function counterPlus() {
        newCounter = counter+1;
        console.log(`saya menambah menjadi : ${newCounter} kali`);
        setCounter(newCounter);
    }

    return (
        <>
            <button onClick={() => clickHandlerParams('Lutfan')}>Klik Nama Saya</button>
            <button onClick={clickHandler}>Klik Saya</button>
            <button onClick={counterPlus}>Counter++ {counter}</button>
        </>
    )
}

export default Button