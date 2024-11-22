import { LitElement, html, css } from "lit";
import { customElement } from "lit/decorators.js";

@customElement("file-search")
class SearchFileComponent extends LitElement {
  searchLetters = "taxi";

  static styles = css`
    @tailwind base;
    @tailwind components;
    @tailwind utilities;

    :host {
      display: block;
      padding: 1rem;
    }
  `;

  render() {
    return html`
      <input
        type="file"
        id="uploaded-file"
        class="file-input file-input-bordered w-full max-w-xs"
      />
      <button class="btn" @click=${this.handleButtonClick}>
        Filter Content
      </button>
    `;
  }

  handleButtonClick() {
    const fileInput = this.shadowRoot?.getElementById(
      "uploaded-file"
    ) as HTMLInputElement;
    console.log(fileInput);
    if (fileInput && fileInput.files && fileInput.files.length > 0) {
      const file = fileInput.files[0];
      const reader = new FileReader();
      reader.onload = (e: any) => {
        const fileContent = e.target.result;
        const updatedFileContent = this.updateFileContent(fileContent);
        this.downloadFile(updatedFileContent);
      };
      reader.readAsText(file);
    }
  }

  updateFileContent(fileContent: string) {
    let updatedFileContent = "";
    fileContent
      .toLocaleLowerCase()
      .split("")
      .forEach((char, _) => {
        if (this.searchLetters.includes(char)) {
          updatedFileContent += char.toUpperCase();
        } else {
          updatedFileContent += char;
        }
      });
    return updatedFileContent;
  }

  downloadFile(fileContent: string) {
    const blob = new Blob([fileContent], { type: "text/plain" });
    const url = URL.createObjectURL(blob);
    const a = document.createElement("a");
    a.setAttribute("href", url);
    a.setAttribute("download", "updated-file.txt");
    a.click();

    URL.revokeObjectURL(url);
  }
}

export default SearchFileComponent;
