import { Component, createSignal, onMount, For, Show } from 'solid-js';
import Identità from './Identità';
import Styles from './App.css';
import Bottoni from './Bottoni'
function App() {

  const [Lista, setLista] = createSignal([]);
  const prendiLista = () => {
    onMount(async () => {
      const res = await fetch(`http://localhost:8080/users`);
      setLista(await res.json());
    });
  }
  
  const changeShow = () => {
    return Lista().length === 0
  }

  return (
    <>
      <Bottoni setLista={setLista} prendiLista={prendiLista}
        idN={() => { return Lista().length + 1 }} class="Style.centro" />
      <Show when={!changeShow()}>
        <table>
          <thead>
            <tr>
              <th>id</th>
              <th>name</th>
              <th>username</th>
              <th>email</th>
              <th>imported</th>
            </tr>
          </thead>
          <tbody>
            <For each={Lista()}>{(user: any, i) =>
              <tr>
                <td>{user.id}</td>
                <td>{user.name}</td>
                <td>{user.username}</td>
                <td>{user.email}</td>
                <td>{String(user.Imported)}</td>
              </tr>
            }</For>
          </tbody>
        </table>
      </Show>
    </>
  )
}

export default App;

/* //const [lista, setLista] = createSignal({nome:"",cognome:""})
  //così non mi funziona
  const [nome, setNome] = createSignal()
 
  return (
    <>
      <button onclick={() => {
        setNome((document.getElementById("nome") as HTMLInputElement).value)
      }}>aggiorna nomi</button>
      <Identità nome={nome()} />
      <Identità nome="gino" cognome="rossi" />
      <Identità nome="paolo" cognome={nome()} />
    </>
  )
}*/