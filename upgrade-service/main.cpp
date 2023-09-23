#include <bits/stdc++.h>

using namespace std;

// Uncomment the following line if you need it
// typedef long long ll;

int main() {
	unsigned long long n=0, m=0, k=0,temp=0,sig=0,temp2=0,i=0;
	cin.sync_with_stdio(false);
	cin >> n >> m >> k;
	unsigned long long *arr=new unsigned long long[n+1];
	for (i = 1; i <= n; i++) {
		unsigned long long ai;
		cin >> ai;
		arr[i]=ai;
	}
    for (i=1;i<=m;i++)
    {
    	temp+=arr[i];
	}
	sig=1;
	for(i=2;i<=n-m+1;)
	{
		temp2=temp+arr[i+m-1]-arr[i-1];
		if (temp2<temp)
		{
			temp=temp2;
			sig=i;
		}
		i++;
	}
	if (temp<(m*k))
	{
		cout << sig <<' '<<sig+m-1 << endl;
	}
	else{
		cout << 0<<' '<< 0 << endl;
	}

	return 0;
}
