import { useContext } from "react"
import { createContext } from 'react'

function Child1(props) {
  return <Child2 name={props.name}/>
}
function Child2(props) {
  return <Child3 name={props.name}/>
}
function Child3(props) {
  const data = useContext(context);
  return <>
    <h1>Halo ini dari child 3 name : {data.name}, umur : {data.umur}67</h1>
  </>
}

const context = createContext();
function App() {
  
  return (
    <>
      <context.Provider value={{name:"Lutfan", umur:20}}>
        <Child1 />
      </context.Provider>
    </>
  )
}

export default App