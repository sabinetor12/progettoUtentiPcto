import { Component, createSignal, onMount, For } from 'solid-js';


export default function Bottoni(props:any) {
    
    const cancella = (event: any) => {
        event.preventDefault()
        props.setLista([])
    }
    const prendiLista = (event: any) => {
        event.preventDefault()
        props.prendiLista();
    }
    const addRow = (event: any) => {
        event.preventDefault()
        props.setLista((Lista: any) => {
            return [...Lista, {
                // id: "1",
                // name: "ciao",
                // username: "squigga",
                // email: "ciao@email.alb",
                // //address: "casa mia"
                id: String(props.idN()),
                name: prompt("ins nome"),
                username: prompt("ins username"),
                email: prompt("ins email"),
                Imported:"false"
            }]
        })
    }

    const scaricaFile = () => {
        location.href ="http://localhost:8080/users/file"      
    }

    return (
        <>
            <button onClick={prendiLista}>prendiLista</button>
            <button onClick={cancella}>cancella</button>
            <button onClick={addRow}>aggiungi</button>
            <button onClick={scaricaFile}>scarica</button>
        </>
    )
}