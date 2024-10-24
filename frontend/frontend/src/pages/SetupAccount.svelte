<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import type { Writable } from "svelte/store";
  import { notifs, TypesType } from "../lib/notifs";
  import { ValidateUsername, CreateUsername } from "../../wailsjs/go/main/App";
  import { user } from "../lib/user";
  import { AuthStatuses } from "../lib/types";
  
  export let authStatus: Writable<AuthStatuses>;
  
  let input: HTMLInputElement;
  let submit: HTMLButtonElement;

  let checkInputTimeout: number;
  let debounceTime = 300;
  let valid: boolean|null = null;
  let invalidReason = "";
  let usernameExists: HTMLSpanElement;
  let stopSVG: SVGElement;

  const switchStatus = () => {
    $authStatus = AuthStatuses.Authenticated;
  }

  onMount(() => {
    const wiggle = (ele: Element) => {
      ele.classList.remove("wiggle");
      ele.getClientRects();
      ele.classList.add("wiggle");
    }
    const validateUsername = () => {
      ValidateUsername($user, input.value)
        .then(result => {
          if (!result.success)
            return notifs.addEndpointError(result);
          valid = result.data.valid;
          if (result.data.reason)
            invalidReason = result.data.reason;
        })
        .catch(error => {
          notifs.add(
            TypesType.Error,
            "Description: " + String(error),
            "Couldn't validate username"
          );
        });
    }
    input.addEventListener("input", () => {
      clearTimeout(checkInputTimeout);
      if (input.value.length === 0)
        valid = null;
      else
        checkInputTimeout = setTimeout(validateUsername, debounceTime);
    });

    submit.addEventListener("click", () => {
      if (!valid) {
        wiggle(usernameExists);
        wiggle(stopSVG);
        return
      }
      CreateUsername($user, input.value)
        .then(result => {
          if (!result.success)
            return notifs.addEndpointError(result);
          $user.username = result.data.username;

          switchStatus();
        })
        .catch(error => {
          notifs.add(
            TypesType.Error,
            "Description: " + String(error),
            "Couldn't create username"
          );
        });
    });
  });
  onDestroy(() => {
    clearTimeout(checkInputTimeout);
  });
</script>
<div class="container">
  <h1 class="hover-color">Lets get you set up!</h1>
  <div>
    <div class="hover-color username-header">
      <h3>Type in your username, make sure it's unique!</h3>
      <!-- svg of a curved arrow -->
      <svg class="curved-arrow" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path fill-rule="evenodd" clip-rule="evenodd" d="M17.5303 13.9697C17.8232 14.2626 17.8232 14.7374 17.5303 15.0303L12.5303 20.0303C12.2374 20.3232 11.7626 20.3232 11.4697 20.0303L6.46967 15.0303C6.17678 14.7374 6.17678 14.2626 6.46967 13.9697C6.76256 13.6768 7.23744 13.6768 7.53033 13.9697L11.25 17.6893L11.25 9.5C11.25 8.78668 11.0298 7.70001 10.3913 6.81323C9.7804 5.96468 8.75556 5.25 7 5.25C6.58579 5.25 6.25 4.91421 6.25 4.5C6.25 4.08579 6.58579 3.75 7 3.75C9.24444 3.75 10.7196 4.70198 11.6087 5.93677C12.4702 7.13332 12.75 8.54665 12.75 9.5L12.75 17.6893L16.4697 13.9697C16.7626 13.6768 17.2374 13.6768 17.5303 13.9697Z" fill="#42210d"/>
      </svg>
    </div>
    <div class="username-cont" class:valid={valid} class:invalid={valid === false}>
      <div class="username">
        <div class="prefix spacing">@</div>
        <div class="divider"></div>
        <input bind:this={input} class="spacing" placeholder="ex. JohnDoe123">
        <div class="checkmark-cont">
          <svg class="checkmark" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd" clip-rule="evenodd" d="M1 12C1 5.92487 5.92487 1 12 1C18.0751 1 23 5.92487 23 12C23 18.0751 18.0751 23 12 23C5.92487 23 1 18.0751 1 12ZM18.4158 9.70405C18.8055 9.31268 18.8041 8.67952 18.4127 8.28984L17.7041 7.58426C17.3127 7.19458 16.6796 7.19594 16.2899 7.58731L10.5183 13.3838L7.19723 10.1089C6.80398 9.72117 6.17083 9.7256 5.78305 10.1189L5.08092 10.8309C4.69314 11.2241 4.69758 11.8573 5.09083 12.2451L9.82912 16.9174C10.221 17.3039 10.8515 17.301 11.2399 16.911L18.4158 9.70405Z" fill="#42210d"/>
          </svg>
        </div>
        <div class="stop-cont">
            <svg bind:this={stopSVG} class="stop" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path fill-rule="evenodd" clip-rule="evenodd" d="M12 22C17.5228 22 22 17.5228 22 12C22 6.47715 17.5228 2 12 2C6.47715 2 2 6.47715 2 12C2 17.5228 6.47715 22 12 22ZM8.58579 8.58579C8 9.17157 8 10.1144 8 12C8 13.8856 8 14.8284 8.58579 15.4142C9.17157 16 10.1144 16 12 16C13.8856 16 14.8284 16 15.4142 15.4142C16 14.8284 16 13.8856 16 12C16 10.1144 16 9.17157 15.4142 8.58579C14.8284 8 13.8856 8 12 8C10.1144 8 9.17157 8 8.58579 8.58579Z" fill="#42210d"/>
            </svg>
        </div>
      </div>
      <div class="username-exists-cont">
        <span bind:this={usernameExists} class="username-exists">{invalidReason}</span>
      </div>
    </div>
  </div>
  <div>
    <div class="hover-color submit-header">
      <h3>And well, that's all. Just click this button</h3>
        <!-- svg of a curved arrow -->
      <svg class="curved-arrow" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path fill-rule="evenodd" clip-rule="evenodd" d="M17.5303 13.9697C17.8232 14.2626 17.8232 14.7374 17.5303 15.0303L12.5303 20.0303C12.2374 20.3232 11.7626 20.3232 11.4697 20.0303L6.46967 15.0303C6.17678 14.7374 6.17678 14.2626 6.46967 13.9697C6.76256 13.6768 7.23744 13.6768 7.53033 13.9697L11.25 17.6893L11.25 9.5C11.25 8.78668 11.0298 7.70001 10.3913 6.81323C9.7804 5.96468 8.75556 5.25 7 5.25C6.58579 5.25 6.25 4.91421 6.25 4.5C6.25 4.08579 6.58579 3.75 7 3.75C9.24444 3.75 10.7196 4.70198 11.6087 5.93677C12.4702 7.13332 12.75 8.54665 12.75 9.5L12.75 17.6893L16.4697 13.9697C16.7626 13.6768 17.2374 13.6768 17.5303 13.9697Z" fill="#42210d"/>
      </svg>
    </div>
    <button bind:this={submit} class="submit">All Done!</button>
  </div>
</div>

<style>
  .submit {
    width: 250px;
    margin: auto;
    cursor: pointer;
    display: block;
    overflow: hidden;
    height: 100%;
    box-sizing: border-box;
    font-size: 30px;
    text-align: center;
    background-color: #705838;
    border: 3px solid #42210d;
    border-radius: 18px;
    color: #42210d;
    font-weight: bold;
    text-align: center;
    transition: border-color .2s, color .2s, background-color .2s, color .2s;
  }
  .submit:active {
    background-color: #42210d;
    color: #705838;
  }
  .submit:hover {
    border-color: #d4bea1;
    color: #d4bea1;
  }
  .username-exists {
    opacity: 0;
    display: block;
    transition: margin-top .25s cubic-bezier(0.25, 0.1, 0.25, 1.4);
    font-weight: bold;
    font-size: 14px;
    margin-top: -28px;
    margin-left: 20px;
  }
  .username-cont:not(.invalid) .username-exists {
    transition: margin-top .25s cubic-bezier(0.25, 0.1, 0.25, 1.4), opacity 0s .3s;
  }
  .username-exists-cont {
    overflow: hidden;
  }
  .username-cont.invalid .username-exists {
    margin-top: 5px;
    opacity: 1;
  }
  .username .checkmark-cont {
    transition: top .25s cubic-bezier(0.25, 0.1, 0.25, 1.4);
    display: flex;
    align-items: center;
    height: 100%;
    position: absolute;
    right: 15px;
    top: 100%;
  }
  .username .checkmark {
    height: 65%;
  }
  .username-cont.valid .checkmark-cont {
    top: 0%;
  }
  .username .stop-cont {
    transition: top .25s cubic-bezier(0.25, 0.1, 0.25, 1.4);
    display: flex;
    align-items: center;
    height: 100%;
    position: absolute;
    right: 15px;
    top: -100%;
  }
  .username .stop {
    height: 65%;
  }
  .username-cont.invalid .stop-cont {
    top: 0%;
  }
  .hover-color {
    transition: color .2s;
  }
  .hover-color:hover {
    color: #d4bea1;
  }
  .username-header svg path, .submit-header svg path {
    transition: fill .2s;
  }
  .username-header:hover svg path, .submit-header:hover svg path {
    fill: #d4bea1;
  }
  .username {
    overflow: hidden;
    height: 100%;
    box-sizing: border-box;
    position: relative;
    display: flex;
    font-size: 25px;
    justify-content: start;
    align-items: center;
    background-color: #705838;
    border: 3px solid #42210d;
    border-radius: 18px;
    transition: border-color .2s;
  }
  .username input, .username .prefix {
    transition: color .2s;
  }
  .username .divider {
    transition: background-color .2s;
  }
  .username-cont:hover .username {
    border-color: #d4bea1;
  }
  .username-cont:hover .divider {
    background-color: #d4bea1;
  }
  /* .username-cont .stop path, .username-cont .checkmark path { */
  /*   transition: fill .2s; */
  /* } */
  /* .username-cont:hover .stop path, .username-cont:hover .checkmark path { */
  /*   fill: #d4bea1; */
  /* } */
  .username-cont {
    margin: auto;
    height: 50px;
    width: 350px;
  }
  .username-cont .username-exists {
    transition: color .2s;
  }
  .username-cont:hover .username-exists {
    color: #d4bea1;
  }
  .username .spacing {
    padding: 0px 10px;
  }
  .username input {
    color: #42210d;
  }
  .username input::placeholder {
    color: inherit;
  }
  .username input:focus, .username:has(input:focus) .prefix, .username input:focus::placeholder {
    fill: #d4bea1;
  }
  .username input {
    font-weight: bold;
    width: 100%;
    background-color: transparent;
    border: none;
    padding: 0px;
    font-size: inherit;
    outline: none;
  }
  .divider {
    height: 100%;
    flex: 0 0 3px;
    background-color: #42210d;
  }
  .prefix {
    /* vertically center the prefix */
    display: flex;
    align-items: center;

    text-align: center;
    height: 100%;
    font-weight: bold;
  }
  .container {
    display: flex;
    align-items: center;
    flex-direction: column;
    gap: 35px;
    margin-top: 110px;
  }
  h1 {
    font-size: 55px;
    margin: 0px;
    text-align: center;
  }
  .submit-header {
    margin-bottom: 10px;
  }
  .username-header, .submit-header {
    display: flex;
    height: 55px;
    justify-content: center;
  }
  .curved-arrow {
    height: 100%;
    transform: rotate(8deg);
  }
</style>
