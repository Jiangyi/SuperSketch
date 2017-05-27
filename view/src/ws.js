class WebsocketConnection {

  init() {
    if (window.WebSocket) {
      this.conn = new WebSocket(`ws://${document.location.host}/ws${document.location.search}`);
      this.conn.onclose = this.onClose;
      this.conn.onmessage = this.onMessage;
    } else {
      console.log('Your browser does not support websockets.');
    }
  }

  static onClose() {
    console.log('Connection closed.');
  }

  static onMessage(event) {
    console.log(event);
  }
}

export default WebsocketConnection;
