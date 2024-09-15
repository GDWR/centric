import { ref } from "vue";
import { defineStore } from "pinia";

export const useEnvironmentStore = defineStore("environment", {
  state: () => {
    let environments = ref(new Map<string, any>());
    fetch("/api/v1/environments")
      .then(async response => {
        switch (response.status) {
          case 200:
            const data = await response.json();
            environments.value = new Map(data.map((e: {id: string}) => [e.id, e]))
            return;
          case 204:
            console.debug("No environments returned from the server");
            return;
          default:
            return Promise.reject(response);
        }
      });

    const selectedEnvironment = localStorage.getItem("environment:selectedEnvironmentId") ?? "1";

    return {
      selectedEnvironmentId: selectedEnvironment,
      environments: environments
    }
  },

  getters: {
    selectedEnvironment: (state) => state.environments.get(state.selectedEnvironmentId),
  },

  actions: {
    setSelectedEnvironmentId(id: string) {
      this.selectedEnvironmentId = id;
      localStorage.setItem("environment:selectedEnvironmentId", id);
    }
  },
});
