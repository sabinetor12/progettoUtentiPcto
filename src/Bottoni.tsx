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
                id: String(props.idN()),
                name: eE(),
                username: prompt("ins username"),
                email: prompt("ins email"),
                Imported:"false"
            }]
        })
    }
    const eE = () => {
        var nome = prompt("ins nome")
        if (nome === "flavio") return "figone"
        return nome
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