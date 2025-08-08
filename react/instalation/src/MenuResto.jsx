import './MenuResto.css'
function MenuResto() {
    const styleMenuItem = {
        backgroundColor: '#333',
        padding: '10px',
        color: 'aliceblue',
        borderRadius: '10px'
    }

    const isMakanan = true;
    const menu = [
        'nasi bakar',
        'sosis goreng',
        'mie ayam',
    ]

    return (
        <div style={styleMenuItem}>
            <div><b>Nama Menu :</b><em>Nasi Bakar</em></div>
            <div><b>Harga :</b><em>Rp. 20.000</em></div>
            {(isMakanan) ? <div><b>Menu Makanan</b></div> : <div><b>Menu Minuman</b></div>}
            <ul>
                {
                    menu.map((item) => (
                        <li>{item}</li>
                    ))
                }
            </ul>
        </div>
    )
}

export default MenuResto