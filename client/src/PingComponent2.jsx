import React, {useState} from "react";
import axios from "axios";

const PingComponent2 = () => {
    const [pingWord, setPingWord] = useState("")
    axios( {
        method: "get",
        url: "api/ping"
    } )
    .then(
        res => {
            console.log(res.data);
            setPingWord(res.data);
        }
    )
    .catch(
        (error) => {
            console.log(error);
    })

    return (
        <h1>
            Ping: {pingWord}
        </h1>
        )
}

export default PingComponent2;