#include <bits/stdc++.h>

using namespace std;

// Uncomment the following line if you need it
// typedef long long ll;

int main() {
	long long n, m, k,temp,sig,sum=0,minsum=0;
	cin.sync_with_stdio(false);
	queue<long long> q1;
	long long q2;
	cin >> n >> m >> k;
	for (long long i = 1; i <= m; i++) {
		long long ai;
		cin >> ai;
		q1.push(ai);
		sum+=ai;    
	}
	temp=1;
	sig=1;
	minsum=sum;
	for (long long j=m+1;j<=n;j++)
	{
		long long bi;
		temp++;
		cin >>bi;
		q1.push(bi);
		sum+=bi-q1.front();
		if (sum<minsum)
		{
			minsum=sum;
			sig=temp;
		}		
		q1.pop();	
	}
	if (minsum < m*k)
	{
		cout << sig<<' '<<sig+m-1<<endl;
	}
    else{
    	cout<<0<<' '<<0<<endl;
	}
	return 0;
}
