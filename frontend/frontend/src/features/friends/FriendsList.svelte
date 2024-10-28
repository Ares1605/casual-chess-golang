<script lang="ts">
  import { onDestroy } from "svelte";
  import { GetFriends } from "../../../wailsjs/go/main/App";
  import { user } from "../../lib/user";
  import { modals } from "../../lib/modals";
  import { notifs, TypesType } from "../../lib/notifs";
  import FriendsCategory from "./FriendsCategory.svelte";
  import FriendListFooter from "./FriendListFooter.svelte";
  import UserItem from "./UserItem.svelte";

  let friends = [];
  GetFriends($user)
    .then(result => {
      if (!result.success)
        return notifs.addEndpointError(result);
      friends = result.data;
    })
    .catch(error => {
      notifs.add(TypesType.Error, String(error), "Fetching friends");
    });

  const minimizedWidth = "60px";

  let expanded = false;

  let closeListTimeout: number;
  const onMouseEnter = () => {
    clearTimeout(closeListTimeout);
  }
  const onMouseLeave = () => {
    clearTimeout(closeListTimeout);
    closeListTimeout = setTimeout(() => expanded = false, 200);
  }
  onDestroy(() => {
    clearTimeout(closeListTimeout);
  });
  
</script>
<div class="container" class:expanded={expanded}>
  <button class="parent-list" on:mouseenter={onMouseEnter} on:mouseleave={onMouseLeave} on:click={() => expanded = true}>
    <div class="list-cont">
      <UserItem {minimizedWidth} user={$user} />
      <FriendsCategory name="ONLINE" users={friends} {minimizedWidth} />
      <FriendsCategory name="OFFLINE" users={friends} {minimizedWidth} />
    </div>
    <div class="footer" class:expanded={expanded}>
      <FriendListFooter />
    </div>
  </button>
  <svg class="arrow-icon" viewBox="90 0 76 256" xmlns="http://www.w3.org/2000/svg">
    <path d="M160,220a11.96287,11.96287,0,0,1-8.48535-3.51465l-80-80a12.00062,12.00062,0,0,1,0-16.9707l80-80a12.0001,12.0001,0,0,1,16.9707,16.9707L96.9707,128l71.51465,71.51465A12,12,0,0,1,160,220Z"/>
  </svg>
</div>

<style lang="scss">
  .footer {
    opacity: 0;
    height: 40px;

    width: 100%;
    visibility: hidden;
    transition: .2s;
    position: absolute;
    bottom: 0px;
    left: 0px;
  }
  .footer.expanded {
    visibility: visible;
    opacity: .9;
  }
  .container {
    display: flex;
    justify-content: right;
    position: relative;
    flex-direction: row-reverse;
    align-items: center;
    height: 100%;
  }
  .arrow-icon {
    opacity: 0;
    transition: .2s opacity;
    width: 9px;
    padding: 5px;
    fill: var(--color-neutral);
  }
  .container:not(.expanded) .parent-list:hover {
    opacity: .93;

    ~ .arrow-icon {
      opacity: 1;
    }
  }
  .container:not(.expanded) .parent-list {
    margin-right: -190px;
  }
  .parent-list {
    position: relative;
    outline: none;
    transition: .2s opacity, .15s margin-right ease-in-out;
    height: 100%;
    box-sizing: border-box;
    min-width: 250px;
    max-width: 250px;
    border: none;
    border-top: 5px solid var(--color-primary-darker);
    background-color: var(--color-primary-darker);
    box-sizing: border-box;
    padding: 0px;
  }
  .list-cont {
    display: flex;
    overflow-y: scroll;
    flex-direction: column;
    padding: 0px 3px 40px 0px;
    height: 100%;

    width: 100%;
    min-height: 100%; /* always grow to the parent element, so the footer sticks to the bottom properly */
    gap: 3px;
    box-sizing: border-box;
  }
  .list-cont::-webkit-scrollbar { 
    display: none;
  }
</style>

