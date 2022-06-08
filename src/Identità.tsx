import { mergeProps } from 'solid-js';

export default function Identità(props:any){
    const merged = mergeProps({ nome: "flavio", cognome: "ndoja" }, props)
    
    return <p> { merged.nome } {merged.cognome} </p> 
    
}