class SessionSummary {
  final String fullSummary;
  final String compressedSummary;

  SessionSummary({required this.fullSummary, required this.compressedSummary});

  factory SessionSummary.fromJson(Map<String, dynamic> json) {
    return SessionSummary(
      fullSummary: json['full_summary'],
      compressedSummary: json['compressed_summary'],
    );
  }
}
