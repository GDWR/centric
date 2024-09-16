import { ref } from "vue";
import { defineStore } from "pinia";
import { useEnvironmentStore } from "./environment";

export const useContainersStore = defineStore("containers", {
  state: () => {
    let containers = ref([]);
    
    const { selectedEnvironmentId } = useEnvironmentStore()
    fetch(`/api/v1/environments/${selectedEnvironmentId}/containers`)
      .then(async response => {
        const data = await response.json();
        containers.value = data;
      });

    return { containers }
  },

  getters: {
    count: (state) => state.containers.length,
  },
});
