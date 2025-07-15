class Message {
  final String sender;
  final String content;

  Message({required this.sender, required this.content});

  Map<String, dynamic> toJson() => {
        'sender': sender,
        'content': content,
      };

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      sender: json['sender'],
      content: json['content'],
    );
  }
}
