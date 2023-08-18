processBtn.addEventListener("click", async () => {
  // Prepare the FormData object to include the selected file
  const formData = new FormData();
  formData.append("excelFile", excelFileInput.files[0]);

  try {
    const response = await fetch("http://localhost:8080/process-excel", {
      method: "POST",
      body: formData,
    });

    const result = await response.json();

    // Clear the existing table rows and header
    tableBody.innerHTML = "";
    const table = document.querySelector("table");
    table.innerHTML = ""; // Clear the table contents

    // Populate the table header with the keys from the JSON response
    const headerRow = document.createElement("tr");
    const headerKeys = Object.keys(result[0].cells);
    headerKeys.forEach((key) => {
      const headerCell = document.createElement("th");
      headerCell.textContent = key;
      headerRow.appendChild(headerCell);
    });
    table.appendChild(document.createElement("thead")).appendChild(headerRow);

    // Populate the table body with the received JSON data
    result.forEach((row) => {
      const rowElement = document.createElement("tr");
      headerKeys.forEach((key) => {
        const cellElement = document.createElement("td");
        cellElement.textContent = row.cells[key];
        rowElement.appendChild(cellElement);
      });
      table.appendChild(document.createElement("tbody")).appendChild(rowElement);
    });
  } catch (error) {
    console.error("Error processing Excel:", error);
  }
});