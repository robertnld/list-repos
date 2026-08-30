"use strict";

const refreshButton = document.querySelector("#refresh-button");

if (refreshButton) {
    refreshButton.addEventListener("click", () => {
        refreshButton.disabled = true;
        refreshButton.textContent = "Refreshing...";

        window.location.reload();
    });
}