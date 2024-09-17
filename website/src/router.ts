import { createRouter, createWebHistory } from "vue-router";

import InitialSetup from "@/views/InitialSetup.vue";
import Dashboard from "./views/Dashboard.vue";
import Containers from "./views/Containers.vue";
import Images from "./views/Images.vue";
import Volumes from "./views/Volumes.vue";
import Networks from "./views/Networks.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: "/", redirect: "/dashboard" },
    { path: "/initial-setup", component: InitialSetup, },
    { path: "/dashboard", component: Dashboard },
    { path: "/containers", component: Containers },
    { path: "/images", component: Images },
    { path: "/volumes", component: Volumes },
    { path: "/networks", component: Networks },
  ],
});


router.beforeEach(async (to, from) => {
  if (to.path === "/initial-setup") {
    return
  }

  const initialSetupComplete = localStorage.getItem("initialSetupComplete");
  if (initialSetupComplete === null || initialSetupComplete === "false") {
    const response = await fetch("/api/v1/system");
    const { initialSetupComplete } = await response.json();
    localStorage.setItem("initialSetupComplete", initialSetupComplete);

    if (initialSetupComplete === false) {
      return { path: "initial-setup" };
    }
  }
});

export default router;
