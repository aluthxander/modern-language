import './MenuResto.css'
function Looping() {
    const styleMenuItem = {
        backgroundColor: '#333',
        padding: '10px',
        color: 'aliceblue',
        borderRadius: '10px',
        marginBottom: '10px'
    }

    const menu = [
        'nasi bakar',
        'sosis goreng',
        'mie ayam',
    ]

    return (
        <>
            {
                menu.map((item, index) => (
                    <div style={styleMenuItem} key={index}>
                        <div><b>Nama Menu :</b><em>{item}</em></div>
                        <div><b>Harga :</b><em>Rp. 20.000</em></div>
                        <div><b>Menu Makanan</b></div>
                    </div>
                ))
            }
        </>
    )
}

export default Looping