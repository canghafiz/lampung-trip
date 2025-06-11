export function formatRupiah(angka) {
  const numberFormat = new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  });
  return numberFormat.format(angka);
}

export function parseRupiahToNumber(rupiahString) {
  if (typeof rupiahString !== "string") {
    return 0;
  }
  let cleanedString = rupiahString
    .replace(/Rp\s?|IDR\s?/gi, "")
    .replace(/\./g, "")
    .replace(/,/g, ".");
  const parsedNumber = parseFloat(cleanedString);
  return isNaN(parsedNumber) ? 0 : parsedNumber;
}
