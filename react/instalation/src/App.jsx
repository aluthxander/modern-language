import MenuResto from "./MenuResto.jsx"
import Header from "./Header.jsx"

function App() {
  const nama = 'Lutfan'
  return (
    <>
      <Header />
      <h1>Hello {nama}</h1>
      <MenuResto />
    </>
  )
}

export default App