const excelFileInput = document.getElementById("excelFileInput");
const processBtn = document.getElementById("processBtn");
const tableBody = document.getElementById("tableBody");
const tableHead = document.getElementById("tableHead");
const loadingDiv = document.getElementById("loading-div");

processBtn.addEventListener("click", async () => {
  // Prepare the FormData object to include the selected file
  const formData = new FormData();
  formData.append("excelFile", excelFileInput.files[0]);

  loadingDiv.style.display = "flex";
  try {
    const response = await fetch("http://localhost:8080/process-excel", {
      method: "POST",
      body: formData,
    });

    const result = await response.json();
    if (result.error) {
      // Clear the existing table rows and header
      tableBody.innerHTML = "";
      tableHead.innerHTML = ""; // Clear the table header
      
      // Display the error
      displayError(result.error);

      

      // Return early to prevent populating the table
      return;
    }
    // Clear the existing table rows and header
    tableBody.innerHTML = "";
    tableHead.innerHTML = ""; // Clear the table header

    const firstRow = result[0];
    const headerKeys = Object.keys(firstRow.cells);
    

    // Populate the table body with the received JSON data
    result.forEach((row, index) => {
      const rowElement = document.createElement("tr");
      headerKeys.forEach((key) => {
        const cellElement = document.createElement(index === 0 ? "th" : "td");
        cellElement.textContent = row.cells[key];
        rowElement.appendChild(cellElement);
      });
      
      if (index === 0) {
        tableHead.appendChild(rowElement);
      } else {
        tableBody.appendChild(rowElement);
      }
    });
  } catch (error) {
    console.error("Error processing Excel:", error);
    displayError(error.message)
  } finally {
    // Hide loading screen
    loadingDiv.style.display = "none";
  }
});

function displayError(message) {
  const errorDiv = document.createElement("div");
  errorDiv.className = "error-message";
  errorDiv.textContent = message;

  const outputDiv = document.getElementById("error-div");
  outputDiv.innerHTML = "";
  outputDiv.appendChild(errorDiv);

  // Remove the error message after 3 seconds
  setTimeout(() => {
    outputDiv.innerHTML = "";
  }, 2000); // 3000 milliseconds = 3 seconds
}