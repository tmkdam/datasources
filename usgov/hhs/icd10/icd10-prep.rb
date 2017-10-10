require 'csv'

ICD10CM_ORDER_FILE = "/Users/untoldone/Downloads/2018-ICD-10-Code-Order-Descriptions/icd10cm_order_2018.txt"

file_contents = File.read ICD10CM_ORDER_FILE

file_lines = file_contents.split("\n")

code_header = [
  "position",
  "code",
  "formatted_code",
  "billable",
  "short_description",
  "long_description"
]

codes = []

def format_icd10_code(code)
  if code.length <= 3
    return code
  else
    return "#{code[0..2]}.#{code[3..-1]}"
  end
end

file_lines.each do |line|
  codes << [
    line[0..5].strip,
    line[6..13].strip,
    format_icd10_code(line[6..13].strip),
    line[14..15].strip,
    line[15..76].strip,
    line[76..-1].strip
  ]
end

CSV.open("icd-10-cm.csv", "w") do |csv|
  csv << code_header
  codes.each do |value|
    csv << value
  end
end