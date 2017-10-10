require 'csv'

CMS_LONG_DX = "/Users/untoldone/working/ICD/ICD-9/ICD-9-CM-v32-master-descriptions/CMS32_DESC_LONG_DX.txt"
CMS_SHORT_DX = "/Users/untoldone/working/ICD/ICD-9/ICD-9-CM-v32-master-descriptions/CMS32_DESC_SHORT_DX.txt"

icd_long = File.read CMS_LONG_DX
icd_short = File.read CMS_SHORT_DX

icd_long_lines = icd_long.split("\n")
icd_short_lines = icd_short.split("\n")

codes = {}

def format_icd9_code(code)
  if code[0] == "E" && code.length <= 4
    return code
  elsif code[0] == "E"
    return "#{code[0..3]}.#{code[4..-1]}"
  elsif code.length <= 3
    return code
  else
    return "#{code[0..2]}.#{code[3..-1]}"
  end
end

icd_long_lines.each do |line|
  code = line[0..5].strip
  desc = line[5..-1].strip

  codes[code] = [
    code,
    format_icd9_code(code),
    desc
  ]
end

icd_short_lines.each do |line|
  code = line[0..5].strip
  desc = line[5..-1].strip

  codes[code] << desc
end

CSV.open("icd-9-cm.csv", "w") do |csv|
  csv << ["code", "formatted_code", "long_description", "short_description"]
  codes.each do |key, value|
    csv << value
  end
end