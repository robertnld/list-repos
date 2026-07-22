"use strict";

const refreshButton = document.querySelector("#refresh-button");

refreshButton.addEventListener("click", () => {
    refreshButton.disabled = true;
    refreshButton.textContent = "Refreshing...";

    window.location.reload();
});