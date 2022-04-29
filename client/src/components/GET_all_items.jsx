import axios from "axios";

const handleGET_items = async() => {
    axios(
        {
            method: "get",
            url:    "api/items"
        }
    )
    .then( res => {
        console.log(res);
        alert( JSON.stringify(res.data) );
    })
}

const GET_all_items = () => {
    return (
        <div>
            <button onClick={handleGET_items}>
                GET_all_items
            </button>
        </div>
    )
}

export default GET_all_items;