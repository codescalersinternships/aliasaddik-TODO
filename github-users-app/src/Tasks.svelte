<script lang="ts">
 
   import axios from 'axios';
import { onMount } from 'svelte';
 

type Task = {
  _id: string;
  title: string;
  done: boolean;
};
let name =  '';

 type GetTasksResponse = {
  
  data: Task[];
};
let todos= {};
$: allTaskPromise = []
onMount(async()=>{
  try {
   const {data, status} = await axios.get ('http://web:9090/');
   console.log(data)
   console.log('response status is: ', status);

  
   allTaskPromise = data ;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.log('error message: ', error.message);
       Error( error.message);
    } else {
      console.log('unexpected error: ', error);
      new Error( 'An unexpected error occurred');
    }}})
    
   const GetTasks =async()=>{
  try {
   const {data, status} = await axios.get ('http://web:9090/');
   console.log(data)
   console.log('response status is: ', status);

  
   allTaskPromise = data ;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.log('error message: ', error.message);
       Error( error.message);
    } else {
      console.log('unexpected error: ', error);
      new Error( 'An unexpected error occurred');
    }}}

const updateDone = async(input) =>{
 
  try {
     let item = JSON.stringify(input)
      const res = await axios.put('http://web:9090', item);
		} catch (err) {
			console.log(err);
		}
  GetTasks();

	};  
  const addTask = async(input) =>{
 
 try {
    let item = JSON.stringify({ done: false, title:input})
     const res = await axios.post('http://web:9090', item);
   } catch (err) {
     console.log(err);
   }
    GetTasks();
 };   

 const deleteTask = async(input) =>{
 
 try {
 
     const res = await axios.delete('http://web:9090/'+ input);
   } catch (err) {
     console.log(err);
   }
    GetTasks();
 };  
  
  
 
</script>
<section>
  <input bind:value={name} type="text"  autocomplete="off" class="input input__lg" placeholder="new task title"  /> 
  <button type="click" class="btn btn-primary" on:click={() => addTask(name)}>Add Task</button>

</section>
 <section>
   
 
      {#each allTaskPromise as task, index(task._id)}
  
        <div>
       
      <li> 
            {index+1}. {task.title}   <input type="checkbox" checked={task.done} on:change={() => {updateDone( {_id: task._id,done: !task.done, title: task.title})}} />
            <button type="click" class="btn btn-danger" on:click={() => deleteTask(task._id)}>Delete Task</button>
          </li>
           
        </div>
      {/each}
  
  </section>  