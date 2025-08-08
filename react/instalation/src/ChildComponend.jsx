import {memo} from 'react'
function ChildComponent(prop) {
    console.log('ini adalah child component rendered');
    
    return (<>
        <h1>Nama Saya : {prop.name}</h1>
        <button onClick={prop.aksi}>Ganti Nama Channel</button>
    </>)
}

export default memo(ChildComponent)