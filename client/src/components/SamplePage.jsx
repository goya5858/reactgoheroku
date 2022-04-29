import axios from "axios";
import GET_all_items from "./GET_all_items";

const SamplePage = () => {
    const handleRoot = async() => {
        axios(  
            {
                method: "get",
                url:    "api"
            }
        )
        .then( res => {
            console.log(res);
            alert(res.data);
        } )
    }
    
    return (
        <div>
            <button onClick={handleRoot}>
                Root
            </button>
            <GET_all_items/>
        </div>
    )
}

export default SamplePage;