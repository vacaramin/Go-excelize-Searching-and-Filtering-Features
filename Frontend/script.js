const excelFileInput = document.getElementById("excelFileInput");
const processBtn = document.getElementById("processBtn");
const tableBody = document.getElementById("tableBody");
const tableHead = document.getElementById("tableHead");

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
    tableHead.innerHTML = ""; // Clear the table header

    const firstRow = result[0];
    const headerKeys = Object.keys(firstRow.cells);
    const headerRow = document.createElement("tr");
    headerKeys.forEach((key) => {
      const headerCell = document.createElement("th");
      headerCell.textContent = key;
      headerRow.appendChild(headerCell);
    });
    tableHead.appendChild(headerRow);

    // Populate the table body with the received JSON data
    result.forEach((row) => {
      const rowElement = document.createElement("tr");
      headerKeys.forEach((key) => {
        const cellElement = document.createElement("td");
        cellElement.textContent = row.cells[key];
        rowElement.appendChild(cellElement);
      });
      tableBody.appendChild(rowElement);
    });
  } catch (error) {
    console.error("Error processing Excel:", error);
  }
});
