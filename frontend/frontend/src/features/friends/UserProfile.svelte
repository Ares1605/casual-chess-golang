<script lang="ts">
  export let user: any;

  let opacity = 0;

  setTimeout(() => opacity = 1);
</script>

<div style="opacity: {opacity}" class="container">
  <img src={user.profile_url} alt="profile">
  {user.username}
</div>
<style>
  img {
    height: 50px;
  }
  .container {
    transition: .2s opacity;
    position: absolute;
    top: 0px;
    left: -10px;
    width: 150px;
    background-color: #d4bea1;
    border: 3px solid #42210d;
    border-radius: 3px;
    padding: 10px;
    transform: translateX(-100%);
  }
</style>
